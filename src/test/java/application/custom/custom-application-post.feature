Feature: Test for creating a custom application

  Background:
    Given url baseUrl
    Given path '/authentication/login'
    And request { username: 'teststudent', password: 'student'}
    When method post
    * def accessToken = response.access_token
    * def stringWith300 = call read('string-generator.js')

  @validCustomApplication
  Scenario: Create a custom application
    Given path '/application/custom'
    And header Authorization = 'Bearer ' + accessToken
    And header Accept = 'application/json'
    And request { content: 'Some content', note: 'Some note'}
    When method post
    Then status 204

  @blankContent
  Scenario: Create a custom application with blank content
    Given path '/application/custom'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { content: '', note: 'Some note'}
    When method post
    Then status 422

  @contentLength
  Scenario: Content length over 255
    Given path '/application/custom'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { content: '#(stringWith300)', note: 'Some note'}
    When method post
    Then status 422

  @blankNote
  Scenario: Create a custom application with blank note
    Given path '/application/custom'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { content: 'Some content', note: ''}
    When method post
    Then status 422

  @noteLength
  Scenario: Note length over 255
    Given path '/application/custom'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { content: 'Some content', note: '#(stringWith300)'}
    When method post
    Then status 422