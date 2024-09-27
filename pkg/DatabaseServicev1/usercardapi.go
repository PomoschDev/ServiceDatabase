package DatabaseServicev1

import "context"

func (s *serverAPI) Cards(ctx context.Context, req *Empty) (*CardsResponse, error) {

	return nil, nil
}

func (s *serverAPI) CreateCard(ctx context.Context, req *CreateCardRequest) (*Card, error) {

	return nil, nil
}

func (s *serverAPI) FindCardById(ctx context.Context, req *FindCardByIdRequest) (*Card, error) {

	return nil, nil
}

func (s *serverAPI) DeleteCardByModel(ctx context.Context, req *Card) (*HTTPCodes, error) {

	return nil, nil
}

func (s *serverAPI) DeleteCardById(ctx context.Context, req *DeleteCardByIdRequest) (*HTTPCodes, error) {

	return nil, nil
}

func (s *serverAPI) UpdateCard(ctx context.Context, req *Card) (*Card, error) {

	return nil, nil
}
