package utils

import (
	"sushee-backend/entity"
	"sushee-backend/httperror"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

func PgConsErrMasker(
	err error,
	consErr entity.ConstraintErrMaskerMap,
	finErr httperror.AppError,
) error {
	if len(consErr) != 0 {
		assertedErr, ok := err.(*pgconn.PgError)
		if !ok {
			return finErr
		}
		if errValue, ok := consErr[assertedErr.ConstraintName]; ok {
			return errValue
		}
		return finErr
	}

	log.Error().Msgf("Unknown error: %s", err.Error())
	return finErr
}
