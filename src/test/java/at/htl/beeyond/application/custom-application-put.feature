Feature: Update a custom application status

  @valid
  Scenario: Update the status of a custom application
    Given url 'http://localhost:8080/custom-application/approve/1'
    And request ''
    When method put
    Then status 204

  @invalid
  Scenario: Test update of non existent custom application
    Given url 'http://localhost:8080/custom-application/approve/9223372036854775807'
    And request ''
    When method put
    Then status 404