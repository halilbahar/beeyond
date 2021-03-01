Feature: Get template endpoint

  Background:
    * url baseUrl
    * path 'template'
    * def insertTemplate = read('classpath:at/htl/beeyond/integration/util/create-template.feature')
    * def insertTemplateResponse = call insertTemplate
    * def template = insertTemplateResponse.template
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Get all templates
    When method GET
    Then status 200
    And assert response.length == 1

  @student
  Scenario: Get all templates as student
    * configure headers = {Authorization: '#(auth(["teacher"]))'}
    * def deleteTemplate = read('classpath:at/htl/beeyond/integration/util/delete-template.feature')
    * def insertTemplateResponse = call deleteTemplate { id: #(template.id) }
    * configure headers = {Authorization: '#(auth(karate.tags))'}
    Given method GET
    Then status 200
    And assert response.length == 0

  @student
  Scenario: Get a deleted template as student
    * configure headers = {Authorization: '#(auth(["teacher"]))'}
    * def deleteTemplate = read('classpath:at/htl/beeyond/integration/util/delete-template.feature')
    * def insertTemplateResponse = call deleteTemplate { id: #(template.id) }
    * configure headers = {Authorization: '#(auth(karate.tags))'}
    Given path template.id
    And method GET
    Then status 404

