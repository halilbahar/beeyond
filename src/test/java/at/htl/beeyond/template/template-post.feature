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
    And request { name: 'Test Template', description: 'This is a test template', content: 'Some yaml %port%', fields: [{wildcard: 'port',label: 'Port of the webserver',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'}]}
    When method post
    Then status 204

  @nameLength
  Scenario: Name length over 255
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: '#(stringWith300)', description: 'This is a test template', content: 'Some yaml', fields: [{wildcard: 'port',label: 'Port of the webserver',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'}] }
    When method post
    Then status 422

  @descriptionLength
  Scenario: Description length over 255
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Test Template', description: '#(stringWith300)', content: 'Some yaml', fields: [{wildcard: 'port',label: 'Port of the webserver',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'}] }
    When method post
    Then status 422

  @blankContent
  Scenario: Create Template with blank content
    Given path '/template'
    And header Authorization = 'Bearer ' + accessToken
    And header Content-Type = 'application/json'
    And request { name: 'Test Template', description: 'This is a test template', content: '', fields: [{wildcard: 'port',label: 'Port of the webserver',description: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit.'}] }
    When method post
    Then status 422