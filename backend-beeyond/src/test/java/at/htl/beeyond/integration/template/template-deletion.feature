Feature: Template deletion endpoint

  Background:
    * url baseUrl + '/template'
    * configure headers = { Authorization: '#(auth(karate.tags))' }
    * def insertTemplate = read('classpath:at/htl/beeyond/integration/util/create-template.feature')
    * def insertTemplateResponse = call insertTemplate
    * def template = insertTemplateResponse.template

  @teacher
  Scenario: Delete a template
    Given path template.id
    When method DELETE
    Then status 204
    Given path template.id
    When method GET
    Then status 200
    And match response.deleted == true

  @teacher
  Scenario: Delete a template that does not exist
    Given path '1000'
    When method DELETE
    Then status 404
