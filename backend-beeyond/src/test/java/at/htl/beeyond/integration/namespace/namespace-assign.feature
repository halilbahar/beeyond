Feature: Namespace assigning endpoint

  Background:
    * url baseUrl + '/namespace'
    * configure headers = { Authorization: '#(auth(karate.tags))' }

  @teacher
  Scenario: Assign 1 user to a valid namespace
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[1]'

  @teacher
  Scenario: Assign 2 user to a valid namespace
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina",
        "moritz"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'

  @teacher
  Scenario: Assign 3 user to a valid namespace
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina",
        "moritz",
        "marc"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[3]'

  @teacher
  Scenario: Assign 2 user to a invalid namespace (wrong characters)
    Given path 'bééyond'
    Given request
    """
    {
      "users": [
        "emina",
        "moritz"
      ]
    }
    """
    When method PUT
    Then status 422
    * print response
    And match response[0].message == 'Not a valid namespace name'

  @teacher
  Scenario: Assign 2 user to a invalid namespace (too long)
    * def generateString = read('string-generator.js')
    Given path generateString()
    Given request
    """
    {
      "users": [
        "emina",
        "moritz"
      ]
    }
    """
    When method PUT
    Then status 422
    * print response
    And match response[0].message == 'This field needs to be between 1 and 50 characters'

  @teacher
  Scenario: Assign 2 user to a invalid namespace (name of user)
    Given path 'marc'
    Given request
    """
    {
      "users": [
        "emina",
        "moritz"
      ]
    }
    """
    When method PUT
    Then status 422
    And match response[0].message == 'Not a valid namespace name'

  @teacher
  Scenario: Assign a nonexistent user to a valid namespace
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
         "test"
      ]
    }
    """
    When method PUT
    Then status 422
    And match response[0].message == 'User with name [test] does not exist'

  @teacher
  Scenario: Assign a user to its own namespace (name of user)
    Given path 'emina'
    Given request
    """
    {
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Then status 422
    And match response[0].message == 'Not a valid namespace name'

  @teacher
  Scenario: Assign 2 user to a valid namespace where one of them already belongs to that namespace
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'

  @teacher
  Scenario: Assign 2 user to a valid namespace where both of them already belong to that namespace
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'

  @teacher
  Scenario: Assign 2 user with the same name (same user) to a valid namespace
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina",
        "emina"
      ]
    }
    """
    When method PUT
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina",
        "emina"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[1]'

  @teacher
  Scenario: Assign 2 user with the same name (same user) to a invalid namespace (name of user)
    Given path 'emina'
    Given request
    """
    {
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Then status 422
    And match response[0].message == 'Not a valid namespace name'

  @teacher
  Scenario: Assign 1 user and remove 1 from a valid namespace
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "marc"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[1]'
    And match response.users[0].name == 'marc'

  @teacher
  Scenario: Remove 1 from a valid namespace (1 user exists already)
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 404

  @teacher
  Scenario: Assign 1 to a valid namespace (1 user exists already)
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina"
      ]
    }
    """
    When method PUT
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'

  @teacher
  Scenario: Replace 2 users with 2 other users (A, B -> C, D)
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "emina",
        "marc"
      ]
    }
    """
    When method PUT
    Given path 'beeyond'
    Given request
    """
    {
      "users": [
        "moritz",
        "stuetz"
      ]
    }
    """
    When method PUT
    And path 'beeyond'
    When method GET
    Then status 200
    And match response.namespace == 'beeyond'
    And match response.users == '#[2]'
    And match response.users contains { 'name': 'moritz', 'id': #number }
    And match response.users contains { 'name': 'stuetz', 'id': #number }
