Feature: Update a custom application status

  @validApprove
  Scenario: Update the status of a custom application
    Given url 'http://localhost:8080/custom-application/approve/1'
    And request ''
    When method put
    Then status 204

  @invalidApprove
  Scenario: Test update of non existent custom application
    Given url 'http://localhost:8080/custom-application/approve/9223372036854775807'
    And request ''
    When method put
    Then status 404

  @validDeny
  Scenario: Update the status of a custom application
    Given url 'http://localhost:8080/custom-application/deny/1'
    And request ''
    When method put
    Then status 204

  @invalidDeny
  Scenario: Test update of non existent custom application
    Given url 'http://localhost:8080/custom-application/deny/9223372036854775807'
    And request ''
    When method put
    Then status 404