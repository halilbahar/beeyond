Feature: Test for creating a template

  Background:
    Given url baseUrl
    Given path '/authentication/login'
    And request { username: 'testteacher', password: 'teacher'}
    When method post
    * def accessToken = response.access_token
    * def stringWith300 = call read('string-generator.js')

  @create
  Scenario: Create a new template
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Test Template', description: 'This is a test template', content: 'Some yaml' }
    When method post
    Then status 204

  @blankContent
  Scenario: Create Template with blank content
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Test Template', description: 'This is a test template', content: '' }
    When method post
    Then status 422

  @nameLength
  Scenario: Name length over 255
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: '#(stringWith300)', description: 'This is a test template', content: 'Some yaml' }
    When method post
    Then status 422

  @descriptionLength
  Scenario: Description length over 255
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Test Template', description: '#(stringWith300)', content: 'Some yaml' }
    When method post
    Then status 422