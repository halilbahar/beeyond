Feature: Get application endpoint

  Background:
    * url baseUrl
    * path 'application'
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Get all applications
    When method GET
    Then status 200
    And match response == '#array'
    # TODO: after clean up is implemented check if the teacher sees all application and the student only his

  @student
  Scenario: Get all applications as student
    When method GET
    Then status 200
    And match response == '#array'
