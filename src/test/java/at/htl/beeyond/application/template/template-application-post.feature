Feature: Create a template application / Upload
  
  Scenario: Create a template application
    Given url 'http://localhost:8080/template'
    And header Content-Type = 'application/json'
    And request { name: "Test Template Application", description: "This is a test template", content: "Some yaml" }
    When method post

    Given url 'http://localhost:8080/template-application'
    And header Content-Type = 'application/json'
    And request { template: { id: 2, name: "Test Template Application", description: "This is a test template", content: "Some yaml"}, user: { name: "it123456" } }
    When method post
    Then status 204


  