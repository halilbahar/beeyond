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
    Given path 'application', 'stop/'+application.id
    When method PATCH
    Then status 200

  @teacher
  Scenario: Finish a not existing application
    Given path 'stop/100'
    When method PATCH
    Then status 404

  @teacher
  Scenario: Finish a pending application
    Given path 'stop/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state RUNNING'

  @teacher
  Scenario: Finish a denied application
    Given path 'deny/'+application.id
    When method PATCH
    Given path 'application', 'stop/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state RUNNING'

  @teacher
  Scenario: Finish a finished application
    Given path 'stop/'+application.id
    When method PATCH
    Given path 'application', 'stop/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state RUNNING'

  @student
  Scenario: Finish a application as a student
    Given path 'stop/'+application.id
    When method PATCH
    Then status 403
