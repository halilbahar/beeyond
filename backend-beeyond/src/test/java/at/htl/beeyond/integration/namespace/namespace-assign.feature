Feature: Namespace assigning endpoint

#  Background:
#    * url baseUrl
#    * path 'namespace'
#    * configure headers = { Authorization: '#(auth(karate.tags))' }

  @teacher
  Scenario: Assign 1 user to a valid namespace

  @teacher
  Scenario: Assign 2 user to a valid namespace

  @teacher
  Scenario: Assign 3 user to a valid namespace

  @teacher
  Scenario: Assign 2 user to a invalid namespace (wrong characters)

  @teacher
  Scenario: Assign 2 user to a invalid namespace (name of user)

  @teacher
  Scenario: Assign a user to its own namespace (name of user)

  @teacher
  Scenario: Assign 2 user to a valid namespace where one of them already belongs to that namespace

  @teacher
  Scenario: Assign 2 user to a valid namespace where both of them already belong to that namespace

  @teacher
  Scenario: Assign 2 user with the same name (same user) to a valid namespace

  @teacher
  Scenario: Assign 2 user with the same name (same user) to a invalid namespace (name of user)

  @teacher
  Scenario: Assign 1 user and remove 1 from a valid namespace

  @teacher
  Scenario: Remove 1 from a valid namespace (1 user exists already)

  @teacher
  Scenario: Assign 1 to a valid namespace (1 user exists already)

  @teacher
  Scenario: Replace 2 users with 2 other users (A, B -> C, D)
