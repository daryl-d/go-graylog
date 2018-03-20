package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSets returns a list of all index sets.
func (client *Client) GetIndexSets(
	skip, limit int,
) ([]graylog.IndexSet, *graylog.IndexSetStats, *ErrorInfo, error) {
	return client.GetIndexSetsContext(context.Background(), skip, limit)
}

// GetIndexSetsContext returns a list of all index sets with a context.
func (client *Client) GetIndexSetsContext(
	ctx context.Context, skip, limit int,
) ([]graylog.IndexSet, *graylog.IndexSetStats, *ErrorInfo, error) {
	indexSets := &graylog.IndexSetsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints.IndexSets, nil, indexSets)
	return indexSets.IndexSets, indexSets.Stats, ei, err
}

// GetIndexSet returns a given index set.
func (client *Client) GetIndexSet(id string) (*graylog.IndexSet, *ErrorInfo, error) {
	return client.GetIndexSetContext(context.Background(), id)
}

// GetIndexSetContext returns a given index set with a context.
func (client *Client) GetIndexSetContext(
	ctx context.Context, id string,
) (*graylog.IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	is := &graylog.IndexSet{}
	ei, err := client.callGet(
		ctx, client.Endpoints.IndexSet(id), nil, is)
	return is, ei, err
}

// CreateIndexSet creates a Index Set.
func (client *Client) CreateIndexSet(indexSet *graylog.IndexSet) (*ErrorInfo, error) {
	return client.CreateIndexSetContext(context.Background(), indexSet)
}

// CreateIndexSetContext creates a Index Set with a context.
func (client *Client) CreateIndexSetContext(
	ctx context.Context, is *graylog.IndexSet,
) (*ErrorInfo, error) {
	if is == nil {
		return nil, fmt.Errorf("IndexSet is nil")
	}
	return client.callPost(ctx, client.Endpoints.IndexSets, is, is)
}

// UpdateIndexSet updates a given Index Set.
func (client *Client) UpdateIndexSet(is *graylog.IndexSet) (*ErrorInfo, error) {
	return client.UpdateIndexSetContext(context.Background(), is)
}

// UpdateIndexSetContext updates a given Index Set with a context.
func (client *Client) UpdateIndexSetContext(
	ctx context.Context, is *graylog.IndexSet,
) (*ErrorInfo, error) {
	if is == nil {
		return nil, fmt.Errorf("IndexSet is nil")
	}
	if is.ID == "" {
		return nil, errors.New("id is empty")
	}
	copiedIndexSet := *is
	copiedIndexSet.ID = ""
	return client.callPut(ctx, client.Endpoints.IndexSet(is.ID), &copiedIndexSet, is)
}

// DeleteIndexSet deletes a given Index Set.
func (client *Client) DeleteIndexSet(id string) (*ErrorInfo, error) {
	return client.DeleteIndexSetContext(context.Background(), id)
}

// DeleteIndexSetContext deletes a given Index Set with a context.
func (client *Client) DeleteIndexSetContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints.IndexSet(id), nil, nil)
}

// SetDefaultIndexSet sets default Index Set.
func (client *Client) SetDefaultIndexSet(id string) (
	*graylog.IndexSet, *ErrorInfo, error,
) {
	return client.SetDefaultIndexSetContext(context.Background(), id)
}

// SetDefaultIndexSetContext sets default Index Set with a context.
func (client *Client) SetDefaultIndexSetContext(
	ctx context.Context, id string,
) (*graylog.IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	is := &graylog.IndexSet{}
	ei, err := client.callPut(ctx, client.Endpoints.SetDefaultIndexSet(id), nil, is)
	return is, ei, err
}
