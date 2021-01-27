Feature: Template application creation endpoint

  Background:
    * url baseUrl
    * path 'application/template'
    * def insertTemplate = read('create-template.feature')
    * call insertTemplate

  @student
  Scenario: Create a valid template application
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        },
        {
          "fieldId": 2,
          "id": 2,
          "value": "8081"
        }
      ],
      "templateId": 1
    }
    """
    When method POST
    Then status 204



  @student
  Scenario: Create a template application with no note
    Given request
    """
    {
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        },
        {
          "fieldId": 2,
          "id": 2,
          "value": "8081"
        }
      ],
      "templateId": 1
    }
    """
    When method POST
    Then status 204

  @student
  Scenario: Create a template application with a too long note
    * def longNote = call read('/at/htl/beeyond/integration:string-generator.js')
    Given request
    """
    {
      "note": "#(longNote)",
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        },
        {
          "fieldId": 2,
          "id": 2,
          "value": "8081"
        }
      ],
      "templateId": 1
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response contains {message: "This field needs to be between 0 and 255 characters",key: "note", value: "#ignore"}
  @student
  Scenario: Create a template application with no templateId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        },
        {
          "fieldId": 2,
          "id": 2,
          "value": "8081"
        }
      ]
    }
    """
    When method POST
    Then status 422
    And match response == [{message:"This field cannot be empty","value":"",key:"templateId"}]

  @student
  Scenario: Create a template application with a non existing templateId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        },
        {
          "fieldId": 2,
          "id": 2,
          "value": "8081"
        }
      ],
      "templateId": 100
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response == [{message: "Template with id 100 does not exist",value: "100",key: "templateId"}]

  @student
  Scenario: Create a template application where at least 1 field is missing
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        }
      ],
      "templateId": 1
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response == [{"message": "Missing field ids: [2], obsolete field ids: []","value": "","key": ""}]

  @student
  Scenario: Create a template application where all fields are provided but one with an invalid fieldId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        },
        {
          "fieldId": 3,
          "id": 2,
          "value": "8081"
        }
      ],
      "templateId": 1
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response == [{"message": "Missing field ids: [2], obsolete field ids: [3]","value": "","key": ""}]

  @student
  Scenario: Create a template application where at least 1 field is provided with no value and fieldId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        },
        {
          "fieldId": 3
        }
      ],
      "templateId": 1
    }
    """
    When method POST
    Then status 422
    #* print response
    And match response contains {"message": "Missing field ids: [2], obsolete field ids: [3]","value": "","key": ""}
    And match response contains {"message": "This field cannot be empty","value": "","key": "value"}

  @student
  Scenario: Create a template with a non existing fieldId
    Given request
    """
    {
      "note": "string",
      "fieldValues": [
        {
          "fieldId": 1,
          "id": 1,
          "value": "4"
        },
        {
          "fieldId": -1,
          "id": 2,
          "value": "8081"
        }
      ],
      "templateId": 1
    }
    """
    When method POST
    Then status 422
    And match response contains {"message":"TemplateField with id -1 does not exist","value":"-1","key":"fieldId"}
