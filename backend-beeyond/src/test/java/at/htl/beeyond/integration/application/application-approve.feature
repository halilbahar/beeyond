Feature: Approve application endpoint

  Background:
    * url baseUrl
    * path 'application'
    * def insertApplication = read('classpath:at/htl/beeyond/integration/util/create-application.feature')
    * def insertApplicationResponse = call insertApplication
    * def application = insertApplicationResponse.application
    * configure headers = {Authorization: '#(auth(karate.tags))', 'Content-Type': 'application/json'}

  @teacher
  Scenario: Approve a valid application
    Given path 'approve/'+application.id
    When method PATCH
    Then status 200

  @teacher
  Scenario: Approve a not existing application
    Given path 'approve/100'
    When method PATCH
    Then status 404

  @teacher
  Scenario: Approve a denied application
    Given path 'deny/'+application.id
    When method PATCH
    Given path 'application', 'approve/'+application.id
    When method PATCH
    Then status 422
    And match response == 'Application is not in state PENDING'

  @teacher
  Scenario: Approve a approved application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application', 'approve/'+application.id
    When method PATCH
    Then status 422
    And match response == 'Application is not in state PENDING'

  @teacher
  Scenario: Approve a stopped application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application','stop/'+application.id
    When method PATCH
    Given path 'application', 'approve/'+application.id
    When method PATCH
    Then status 422
    And match response == 'Application is not in state PENDING'

  @teacher
  Scenario: Approve a finished application
    Given path 'approve/'+application.id
    When method PATCH
    Given path 'application', 'finish/'+application.id
    When method PATCH
    Given path 'application', 'approve/'+application.id
    When method PATCH
    Then status 422
    And match response == 'Application is not in state PENDING'

  @student
  Scenario: Approve a application as a student
    Given path 'approve/'+application.id
    When method PATCH
    Then status 403
