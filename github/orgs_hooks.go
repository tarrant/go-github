// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import "fmt"

const orgHookPreviewHeader = "application/vnd.github.sersi-preview+json"

// CreateHook creates a Hook for the specified organization.
// Name and Config are required fields.
//
// GitHub API docs: http://developer.github.com/v3/orgs/hooks/#create-a-hook
func (s *OrganizationsService) CreateHook(owner string, hook *Hook) (*Hook, *Response, error) {
	u := fmt.Sprintf("orgs/%v/hooks", owner)
	req, err := s.client.NewRequest("POST", u, hook)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", orgHookPreviewHeader)

	h := new(Hook)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

// ListHooks lists all Hooks for the specified organization.
//
// GitHub API docs: http://developer.github.com/v3/orgs/hooks/#list
func (s *OrganizationsService) ListHooks(owner string, opt *ListOptions) ([]Hook, *Response, error) {
	u := fmt.Sprintf("orgs/%v/hooks", owner)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", orgHookPreviewHeader)

	hooks := new([]Hook)
	resp, err := s.client.Do(req, hooks)
	if err != nil {
		return nil, resp, err
	}

	return *hooks, resp, err
}

// GetHook returns a single specified Hook.
//
// GitHub API docs: http://developer.github.com/v3/orgs/hooks/#get-single-hook
func (s *OrganizationsService) GetHook(owner string, id int) (*Hook, *Response, error) {
	u := fmt.Sprintf("orgs/%v/hooks/%d", owner, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", orgHookPreviewHeader)
	hook := new(Hook)
	resp, err := s.client.Do(req, hook)
	return hook, resp, err
}

// EditHook updates a specified Hook.
//
// GitHub API docs: http://developer.github.com/v3/orgs/hooks/#edit-a-hook
func (s *OrganizationsService) EditHook(owner string, id int, hook *Hook) (*Hook, *Response, error) {
	u := fmt.Sprintf("orgs/%v/hooks/%d", owner, id)
	req, err := s.client.NewRequest("PATCH", u, hook)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", orgHookPreviewHeader)
	h := new(Hook)
	resp, err := s.client.Do(req, h)
	return h, resp, err
}

// DeleteHook deletes a specified Hook.
//
// GitHub API docs: http://developer.github.com/v3/orgs/hooks/#delete-a-hook
func (s *OrganizationsService) DeleteHook(owner string, id int) (*Response, error) {
	u := fmt.Sprintf("orgs/%v/hooks/%d", owner, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", orgHookPreviewHeader)
	return s.client.Do(req, nil)
}

// TestHook triggers a test Hook by github.
//
// GitHub API docs: http://developer.github.com/v3/orgs/hooks/#test-a-push-hook
func (s *OrganizationsService) TestHook(owner string, id int) (*Response, error) {
	u := fmt.Sprintf("orgs/%v/hooks/%d/tests", owner, id)
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", orgHookPreviewHeader)
	return s.client.Do(req, nil)
}
