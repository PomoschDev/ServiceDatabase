package DatabaseServicev1

import "context"

func (s *serverAPI) Wards(ctx context.Context, req *Empty) (*WardsResponse, error) {

	return nil, nil
}

func (s *serverAPI) CreateWard(ctx context.Context, req *CreateWardRequest) (*Ward, error) {

	return nil, nil
}

func (s *serverAPI) FindWardById(ctx context.Context, req *FindWardByIdRequest) (*Ward, error) {

	return nil, nil
}

func (s *serverAPI) DeleteWardByModel(ctx context.Context, req *Ward) (*HTTPCodes, error) {

	return nil, nil
}

func (s *serverAPI) DeleteWardById(ctx context.Context, req *DeleteWardByIdRequest) (*HTTPCodes, error) {

	return nil, nil
}

func (s *serverAPI) UpdateWard(ctx context.Context, req *Ward) (*Ward, error) {

	return nil, nil
}
