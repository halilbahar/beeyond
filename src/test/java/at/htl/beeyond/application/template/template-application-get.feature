Feature: 'GET' tests for TemplateApplication

  Background:
    Given url 'http://localhost:8080/authentication/login'
    And request { username: 'it123456', password: 'passme'}
    When method post
    * def accessToken = response.access_token

  @all
  Scenario: Get list with all template applications
    Given url 'http://localhost:8080/template-application'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    When method get
    Then status 200

  @validId
  Scenario: Get a template application by id
    Given url 'http://localhost:8080/template-application/2'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    When method get
    Then status 200
    And match response.user.name == "it123456"

  @invalidId
  Scenario: Get a not existing template application by id
    Given url 'http://localhost:8080/template-application/9999999999'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    When method get
    Then status 404