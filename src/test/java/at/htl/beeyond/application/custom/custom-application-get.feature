Feature: Test for getting all custom application

  Scenario: Get list with all custom applications
    Given url 'http://localhost:8080/authentication/login'
    And request { username: 'it123456', password: 'passme'}
    When method post
    Then status 200
    * def accessToken = response.access_token

    Given url 'http://localhost:8080/custom-application'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    When method get
    Then status 200
