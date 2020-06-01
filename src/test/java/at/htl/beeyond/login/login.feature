Feature: Test a user login

  Scenario: User wants to login
    Given url 'http://localhost:8080/authentication/login'
    And request { username: 'it123456', password: 'passme'}
    When method post
    Then status 200
    And print response