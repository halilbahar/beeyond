Feature: Get template endpoint

  Background:
    * url baseUrl
    * path 'template'
    * def insertMethod = read('create-two-templates.feature')
    * callonce insertMethod
    * def responseMessage = read('expected-get-response.json')

  @teacher
  Scenario: Get all templates
    When method GET
    Then status 200
    And assert response.length > 0

  @student
  Scenario: Get all templates as student
    When method GET
    Then status 200
    And assert response.length > 0
