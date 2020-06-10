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
    Given path '/application/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request {templateId: 1, name: 'Nginx 1.16',description: 'This is a simple template for a nginx server', content: 'Some content', fields: [{wildcard: 'port',label: '#(stringWith300)',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'}], note: 'Some note'}
    When method post
    Then status 422

  @wildcardLength
  Scenario: Wildcard length over 255
    Given path '/application/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request {templateId: 1, name: 'Nginx 1.16',description: 'This is a simple template for a nginx server', content: 'Some content', fields: [{wildcard: '#(stringWith300)',label: 'Port of the webserver',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'}], note: 'Some note'}
    When method post
    Then status 422