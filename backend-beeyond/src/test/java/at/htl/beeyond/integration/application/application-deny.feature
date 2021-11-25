Feature: Deny application endpoint

  Background:
    * url baseUrl
    * path 'application'
    * def insertApplication = read('classpath:at/htl/beeyond/integration/util/create-application.feature')
    * def insertApplicationResponse = call insertApplication
    * def application = insertApplicationResponse.application
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Deny a valid application
    Given path 'deny/'+application.id
    When method PATCH
    Then status 200

  @teacher
  Scenario: Deny a not existing application
    Given path 'deny/100'
    When method PATCH
    Then status 404

  @teacher
  Scenario: Deny a approved application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application', 'deny/'+application.id
    When method PATCH
    Then status 422
    And match response == 'Application is not in state PENDING'

  @teacher
  Scenario: Deny a denied application
    Given path 'deny/'+application.id
    When method PATCH
    Given path 'application','deny/'+application.id
    When method PATCH
    Then status 422
    And match response == 'Application is not in state PENDING'

  @student
  Scenario: Deny a application as a student
    Given path 'deny/'+application.id
    When method PATCH
    Then status 403
