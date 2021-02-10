Feature: Get template endpoint

  Background:
    * url baseUrl
    * path 'template'
    * def insertMethod = read('classpath:at/htl/beeyond/integration/util/create-template.feature')
    * callonce insertMethod

  @teacher
  Scenario: Get all templates
    When method GET
    Then status 200
    And assert response.length > 0
    # TODO: after a global clean up is programmed replace these length with actual tests

  @student
  Scenario: Get all templates as student
    When method GET
    Then status 200
    And assert response.length > 0
