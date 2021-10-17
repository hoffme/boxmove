package app

// func (s *App) NewBox(client string, params *box.CreateParams) (*box.Box, error) {
// 	return s.Boxes.New(client, params)
// }

// func (s *App) NewMove(client string, params *move.CreateParams) (*move.Move, error) {
// 	extremes, err := s.Boxes.FindAll(client, &box.Filter{ID: []string{params.From, params.To}})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(extremes) != 2 {
// 		return nil, errors.New("invalid extremes from move")
// 	}

// 	boxFrom := extremes[0]
// 	boxTo := extremes[1]

// 	if boxFrom.Id() != params.From {
// 		boxTo = extremes[1]
// 		boxFrom = extremes[0]
// 	}

// 	movement, err := s.Moves.New(client, params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	errBoxFrom := boxFrom.AddActive(params.Active, int64(params.Count)*-1)
// 	if errBoxFrom != nil {
// 		log.Println(errBoxFrom)
// 	}

// 	errBoxTo := boxTo.AddActive(params.Active, int64(params.Count))
// 	if errBoxTo != nil {
// 		log.Println(errBoxTo)
// 	}

// 	return movement, nil
// }
