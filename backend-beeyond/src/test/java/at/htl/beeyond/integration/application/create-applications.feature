Feature:
  Background:
    * url baseUrl
    * path 'application/custom'
    * def content = read('../template/nginx-deployment-template.yml.txt')

  @teacher
  Scenario: Create a valid application
    * configure headers = { Authorization: '#(teacherAuth())'}
    Given request
    """
    {
      "note": "Nginx deployment done by teacher",
      "content": "#(content)",
    }
    """
    When method POST
    Then status 204

  @student
  Scenario: Create a valid as student application
    * configure headers = { Authorization: '#(studentAuth())'}
    Given request
    """
    {
      "note": "Nginx deployment done by student",
      "content": "#(content)",
    }
    """
    When method POST
    Then status 204
