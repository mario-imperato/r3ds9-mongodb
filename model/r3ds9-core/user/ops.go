package user

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func FindByNickname(collection *mongo.Collection, sid string, mustFind bool, findOptions *options.FindOneOptions) (*User, error) {

	const SemLogContext = "r3ds9-core/user/find-by-nickname"
	log.Trace().Str("nickname", sid).Msg(SemLogContext)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := User{}

	f := Filter{}
	f.Or().AndNicknameEqTo(sid)
	err := collection.FindOne(ctx, f.Build(), findOptions).Decode(&ent)
	if err != nil && (err != mongo.ErrNoDocuments || (err == mongo.ErrNoDocuments && mustFind)) {
		log.Error().Err(err).Msg(SemLogContext)
		return nil, err
	} else {
		if err != nil {
			log.Trace().Str("sid", sid).Msgf("%s - document not found", SemLogContext)
			return nil, err
		} else {
			log.Trace().Str("sid", sid).Msgf("%s - document found", SemLogContext)
		}
	}

	return &ent, nil
}

/*
func Insert(ctx context.Context, aCollection *mongo.Collection, s *Session) (string, error) {

	const SemLogContext = "r3ds9-core/session/insert"
	log.Trace().Str("nickname", s.Nickname).Msg(SemLogContext)

	now := time.Now()
	s.Sysinfo.Status = "active"
	s.Sysinfo.Createdat = primitive.NewDateTimeFromTime(now)
	s.Sysinfo.Modifiedat = primitive.NewDateTimeFromTime(now)
	r, err := aCollection.InsertOne(ctx, s)

	if err != nil {
		return "", err
	}

	s.OId = r.InsertedID.(primitive.ObjectID)
	sid := s.SessionId()
	log.Trace().Str("sid", sid).Msg(SemLogContext)

	return sid, nil
}

func Remove(ctx context.Context, aCollection *mongo.Collection, sid string, mustDelete bool) (int, error) {

	const SemLogContext = "r3ds9-core/session/remove"
	log.Trace().Str("sid", sid).Msg(SemLogContext)

	f := Filter{}
	f.Or().AndSessionIdEqTo(sid)
	var opts []*options.DeleteOptions
	r, err := aCollection.DeleteOne(ctx, f.Build(), opts...)
	if err != nil {
		return 0, err
	}

	if mustDelete && r.DeletedCount != 1 {
		return int(r.DeletedCount), fmt.Errorf("%s - number of affected rows mismatch: actual=%d, wanted:%d", SemLogContext, 1, r.DeletedCount)
	}

	return int(r.DeletedCount), nil
}

func UpdateSession(ctx context.Context, aCollection *mongo.Collection, sid string, mustMatch bool, opts ...UpdateOption) (int, error) {

	const SemLogContext = "r3ds9-core/session/remove"
	log.Trace().Str("sid", sid).Msg(SemLogContext)

	ud := UpdateDocument{}
	for _, o := range opts {
		o(&ud)
	}

	f := Filter{}
	f.Or().AndSessionIdEqTo(sid)
	var uopts []*options.UpdateOptions
	r, err := aCollection.UpdateOne(ctx, f.Build(), ud.Build(), uopts...)
	if err != nil {
		return 0, err
	}

	if mustMatch && r.MatchedCount != 1 {
		return int(r.MatchedCount), fmt.Errorf("%s - number of affected rows mismatch: actual=%d, wanted:%d", SemLogContext, 1, r.MatchedCount)
	}

	return int(r.MatchedCount), nil
}
*/
