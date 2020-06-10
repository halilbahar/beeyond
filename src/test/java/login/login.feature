Feature: Test a user login

  Scenario: User wants to login
    Given url baseUrl + '/authentication/login'
    And request { username: 'teststudent', password: 'student'}
    When method post
    Then status 200
    And print response