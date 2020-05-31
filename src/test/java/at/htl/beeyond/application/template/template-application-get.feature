Feature: 'GET' tests for TemplateApplication

  @all
  Scenario: Get list with all template applications
    Given url 'http://localhost:8080/template-application'
    And header Accept = 'application/json'
    When method get
    Then status 200

  @validId
  Scenario: Get a template application by id
    Given url 'http://localhost:8080/template-application/2'
    And header Accept = 'application/json'
    When method get
    Then status 200
    And match response == { template: { id: 2, name: "Test Template Application", description: "This is a test template", content: "Some yaml"}, user: { name: "it123456" } }

  @invalidId
  Scenario: Get a not existing template application by id
    Given url 'http://localhost:8080/template-application/9999999999'
    And header Accept = 'application/json'
    When method get
    Then status 404