package historydb

import (
	"database/sql"
	"fmt"

	"github.com/tokamak-network/syb-sequencer/sequencer/common"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

func (hdb *HistoryDB) AddVouch(vouch *common.Vouch) error {
	query := `
		INSERT INTO vouch (idx, from_idx, from_eth_addr, to_idx, to_eth_addr)
		VALUES ($1, $2, $3, $4, $5);
	`
	_, err := hdb.dbWrite.Exec(
		query,
		vouch.Idx,
		vouch.FromIdx,
		vouch.FromEthAddr.Bytes(),
		vouch.ToIdx,
		vouch.ToEthAddr.Bytes(),
	)

	if err != nil {
		return fmt.Errorf("AddVouch: failed to add vouch with idx %d: %w", vouch.Idx, err)
	}
	return nil
}

func (hdb *HistoryDB) GetVouchByIdx(idx common.VouchIdx) (*common.Vouch, error) {
	query := `
		SELECT from_idx, from_eth_addr, to_idx, to_eth_addr
		FROM vouch
		WHERE idx = $1;
	`
	row := hdb.dbRead.QueryRow(query, idx)

	v := &common.Vouch{Idx: idx}
	var fromEthAddrBytes, toEthAddrBytes []byte

	err := row.Scan(
		&v.FromIdx,
		&fromEthAddrBytes,
		&v.ToIdx,
		&toEthAddrBytes,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetVouchByIdx: failed to get vouch with idx %d: %w", idx, err)
	}

	v.FromEthAddr = ethCommon.BytesToAddress(fromEthAddrBytes)
	v.ToEthAddr = ethCommon.BytesToAddress(toEthAddrBytes)

	return v, nil
}

func (hdb *HistoryDB) GetVouchesByEthAddress(ethAddr ethCommon.Address) ([]*common.Vouch, error) {
	query := `
		SELECT idx, from_idx, from_eth_addr, to_idx, to_eth_addr
		FROM vouch
		WHERE from_eth_addr = $1;
	`
	rows, err := hdb.dbRead.Query(query, ethAddr.Bytes())
	if err != nil {
		return nil, fmt.Errorf("GetVouchesByEthAddress: failed to query vouches for eth_addr %s: %w", ethAddr.String(), err)
	}
	defer rows.Close()

	var vouches []*common.Vouch
	for rows.Next() {
		v := &common.Vouch{}
		var fromEthAddrBytes, toEthAddrBytes []byte
		errScan := rows.Scan(
			&v.Idx,
			&v.FromIdx,
			&fromEthAddrBytes,
			&v.ToIdx,
			&toEthAddrBytes,
		)
		if errScan != nil {
			return nil, fmt.Errorf("GetVouchesByEthAddress: failed to scan vouch row for eth_addr %s: %w", ethAddr.String(), errScan)
		}
		v.FromEthAddr = ethCommon.BytesToAddress(fromEthAddrBytes)
		v.ToEthAddr = ethCommon.BytesToAddress(toEthAddrBytes)
		vouches = append(vouches, v)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("GetVouchesByEthAddress: error iterating vouch rows for eth_addr %s: %w", ethAddr.String(), err)
	}

	return vouches, nil
}

func (hdb *HistoryDB) GetVouchersByIdx(toIdx common.AccountIdx) ([]common.AccountIdx, error) {
	query := `
        SELECT DISTINCT from_idx 
        FROM vouch 
        WHERE to_idx = $1
        ORDER BY from_idx;
    `

	rows, err := hdb.dbRead.Query(query, toIdx)
	if err != nil {
		return nil, fmt.Errorf("GetVouchersByToIdx: failed to query vouchers: %w", err)
	}
	defer rows.Close()

	var vouchers []common.AccountIdx
	for rows.Next() {
		var fromIdx common.AccountIdx
		err := rows.Scan(&fromIdx)
		if err != nil {
			return nil, fmt.Errorf("GetVouchersByToIdx: failed to scan from_idx: %w", err)
		}
		vouchers = append(vouchers, fromIdx)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("GetVouchersByToIdx: error iterating rows: %w", err)
	}

	return vouchers, nil
}

func (hdb *HistoryDB) DeleteVouchByIdx(idx common.VouchIdx) error {
	query := `DELETE FROM vouch WHERE idx = $1;`
	result, err := hdb.dbWrite.Exec(query, idx)
	if err != nil {
		return fmt.Errorf("DeleteVouchByIdx: failed to delete vouch with idx %d: %w", idx, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteVouchByIdx: failed to get rows affected for vouch idx %d: %w", idx, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("DeleteVouchByIdx: no vouch found with idx %d to delete", idx)
	}
	return nil
}
