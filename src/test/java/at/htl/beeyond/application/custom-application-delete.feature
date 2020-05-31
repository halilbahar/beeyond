Feature: Test for deleting a custom application

  @valid
  Scenario: Delete a custom application
    Given url 'http://localhost:8080/custom-application/1'
    And header Accept = 'application/json'
    When method delete
    Then status 200
    And match response == { id: 1, content:'Some yaml', note: 'Some note', status: 'PENDING'}

  @invalid
  Scenario: Delete a custom application

    Given url 'http://localhost:8080/custom-application/9999999999'
    And header Accept = 'application/json'
    When method delete
    Then status 404


