Feature: Stop application endpoint

  Background:
    * url baseUrl
    * path 'application'
    * def insertApplication = read('classpath:at/htl/beeyond/integration/util/create-application.feature')
    * def insertApplicationResponse = call insertApplication
    * def application = insertApplicationResponse.application
    * def insertRunningApplication = read('classpath:at/htl/beeyond/integration/util/create-student-application.feature')
    * def insertRunningApplicationResponse = call insertRunningApplication
    * def id = insertRunningApplicationResponse.runningApplication.id
    * def approveApplication = read('classpath:at/htl/beeyond/integration/util/approve-student-application.feature')
    * call approveApplication {id: '#(id)'}
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Stop a valid application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application', 'stop/'+application.id
    When method PATCH
    Then print karate.prevRequest.headers
    Then status 200

  @teacher
  Scenario: Stop a not existing application
    Given path 'stop/100'
    When method PATCH
    Then status 404

  @teacher
  Scenario: Stop a pending application
    Given path 'stop/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state RUNNING'

  @teacher
  Scenario: Stop a denied application
    Given path 'deny/'+application.id
    When method PATCH
    Given path 'application', 'stop/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state RUNNING'

  @teacher
  Scenario: Stop a finished application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application','finish/'+application.id
    When method PATCH
    Given path 'application', 'stop/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state RUNNING'

  @teacher
  Scenario: Stop a stopped application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application','stop/'+application.id
    When method PATCH
    Given path 'application', 'stop/'+application.id
    When method PATCH
    Then status 422
    And match response == 'Application is not in state RUNNING'

  @student
  Scenario: Stop an application from someone else as a student
    Given path 'stop/'+application.id
    When method PATCH
    Then status 403

  @student
  Scenario: Stop an application as a student
    Given path 'stop/'+id
    When method PATCH
    Then status 200

