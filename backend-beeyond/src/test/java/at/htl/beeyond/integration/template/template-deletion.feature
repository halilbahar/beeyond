Feature: Template deletion endpoint

  Background:
    * url baseUrl
    * path 'template'
    * def insertMethod = read('create-two-templates.feature')

  @teacher
  Scenario: Delete a template
    Given path '1'
    When method DELETE
    Then status 204

  @teacher
  Scenario: Check if property deleted was actually set to true
    * def fun = function(x) { return x.id == 1 }
    When method GET
    Then status 200
    And match karate.filter(response, fun)[0].deleted == true

  @teacher
  Scenario: Delete a template - template does not exist
    Given path '-1'
    When method DELETE
    Then status 404
