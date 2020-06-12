Feature: Test for invalid template fields

  Background:
    Given url baseUrl
    Given path '/authentication/login'
    And request { username: 'testteacher', password: 'teacher'}
    When method post
    * def accessToken = response.access_token
    * def stringWith300 = call read('string-generator.js')

  @labelLength
  Scenario: Label length over 255
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Test Template', description: 'This is a test template', content: 'Some yaml %port%', fields: [{wildcard: 'port',label: '#(stringWith300)',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'}]}
    When method post
    Then status 422

  @wildcardLength
  Scenario: Wildcard length over 255
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Test Template', description: 'This is a test template', content: 'Some yaml %#(stringWith300)%', fields: [{wildcard: '#(stringWith300)',label: 'Port of the webserver',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'}]}
    When method post
    Then status 422

  @descriptionLength
  Scenario: Description length over 255
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Test Template', description: 'This is a test template', content: 'Some yaml %port%', fields: [{wildcard: 'port',label: 'Port of the webserver',description: '#(stringWith300)'}]}
    When method post
    Then status 422
    And match response[0].message == 'Missing fields: [], obsolete fields: [port]'