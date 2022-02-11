Feature: Start application endpoint

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
    * def stopApplication = read('classpath:at/htl/beeyond/integration/util/stop-student-application.feature')
    * call stopApplication {id: '#(id)'}
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Start a valid application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application', 'stop/'+application.id
    When method PATCH
    Given path 'application', 'start/'+application.id
    When method PATCH
    Then status 200

  @teacher
  Scenario: Start a not existing application
    Given path 'start/100'
    When method PATCH
    Then status 404

  @teacher
  Scenario: Start a pending application
    Given path 'start/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state STOPPED'

  @teacher
  Scenario: Start a denied application
    Given path 'deny/'+application.id
    When method PATCH
    Given path 'application', 'start/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state STOPPED'

  @teacher
  Scenario: Start a finished application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application','finish/'+application.id
    When method PATCH
    Given path 'application', 'start/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state STOPPED'

  @student
  Scenario: Start an application from someone else as a student
    Given path 'start/'+application.id
    When method PATCH
    Then status 403

  @student
  Scenario: Start an application as a student
    Given path 'start/'+id
    When method PATCH
    Then status 200
