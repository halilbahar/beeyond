Feature: Request application endpoint

  Background:
    * url baseUrl
    * path 'application'
    * def insertApplication = read('classpath:at/htl/beeyond/integration/util/create-application.feature')
    * def insertApplicationResponse = call insertApplication
    * def application = insertApplicationResponse.application
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Request a valid application
    Given path 'deny/'+application.id
    When method PATCH
    Given path 'application', 'request/'+application.id
    When method PATCH
    Then status 200

  @teacher
  Scenario: Request a not existing application
    Given path 'request/100'
    When method PATCH
    Then status 404

  @teacher
  Scenario: Request a pending application
    Given path 'request/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state DENIED'

  @teacher
  Scenario: Request a approved application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application', 'request/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state DENIED'

  @teacher
  Scenario: Request a finished application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application','finish/'+application.id
    When method PATCH
    Given path 'application', 'request/'+application.id
    When method PATCH
    Then status 422
    And match response  == 'Application is not in state DENIED'

  @teacher
  Scenario: Request a stopped application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application','stop/'+application.id
    When method PATCH
    Given path 'application', 'request/'+application.id
    When method PATCH
    Then status 422
    And match response == 'Application is not in state DENIED'