Feature: Finish application endpoint

  Background:
    * url baseUrl
    * path 'application'
    * def insertApplication = read('classpath:at/htl/beeyond/integration/util/create-application.feature')
    * def insertApplicationResponse = call insertApplication
    * def application = insertApplicationResponse.application
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Finish a valid application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application', 'finish/'+application.id
    When method PATCH
    Then status 200

  @teacher
  Scenario: Finish a not existing application
    Given path 'finish/100'
    When method PATCH
    Then status 404

  @teacher
  Scenario: Finish a pending application
    Given path 'finish/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state RUNNING or STOPPED'

  @teacher
  Scenario: Finish a denied application
    Given path 'deny/'+application.id
    When method PATCH
    Given path 'application', 'finish/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state RUNNING or STOPPED'

  @teacher
  Scenario: Finish a finished application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application','finish/'+application.id
    When method PATCH
    Given path 'application', 'finish/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state RUNNING or STOPPED'

  @teacher
  Scenario: Finish a stopped application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application','stop/'+application.id
    When method PATCH
    Given path 'application', 'finish/'+application.id
    When method PATCH
    Then status 200

  @student
  Scenario: Finish a application as a student
    Given path 'finish/'+application.id
    When method PATCH
    Then status 403
