Feature: Test for creating a template

  Scenario: Create a new template
    Given url 'http://localhost:8080/authentication/login'
    And request { username: 'it123456', password: 'passme'}
    When method post
    * def accessToken = response.access_token

    Given url 'http://localhost:8080/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: "Test Template", description: "This is a test template", content: "Some yaml" }
    When method post
    Then status 204