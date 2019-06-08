package jira

type ChildrenService struct {
    client *Client
}

type Child struct {
    ID      string `json:"id,omitempty" structs:"id,omitempty"`
    Key     string `json:"key,omitempty" structs:"key,omitempty"`
    Summary string `json:"summary,omitempty" structs:"summary,omitempty"`
}

func (s *ChildrenService) GetImmediateChildren() ([]Child, *response, error) {
    apiEndoint := "rest/synapse/latest/public/"
    req, err := s.client.NewRequest("GET", apiEndpoint, nil)

    if err != nil {
        return nil, nil, err
    }

    ChildrenList := []Child{}

    resp, err := s.client.Do(req, &ChildrenList)

    if err != nil {
        return nil, resp, NewJiraError(resp, err)
    }
    return ChildrenList, resp, nil
}


