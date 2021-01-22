Feature: Get application endpoint

  Background:
    * url baseUrl
    * path 'application'
    * def responseMessageTeacher = read('expected-get-response-teacher.json')
    * def responseMessageStudent = read('expected-get-response-student.json')

  @teacher
  Scenario: Get all applications
    When method GET
    Then status 200
    And match response contains responseMessageTeacher
    And match response contains responseMessageStudent

  @student
  Scenario: Get all applications as student
    When method GET
    Then status 200
    And match response contains responseMessageStudent
    And match response !contains responseMessageTeacher

