package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/amyunfei/simplebank/util"
	"github.com/stretchr/testify/require"
)

func cerateRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	// 报错及存储校验
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	// 数据准确性校验
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	// 自动创建数值校验
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreateAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	cerateRandomEntry(t)
}

func TestGetEntry(t *testing.T) { // 根据id获取账户记录数据 测试
	entry1 := cerateRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	// 数据基本校验
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	// 数据准确性校验
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.CreateAt, entry2.CreateAt)
}

func TestListEntry(t *testing.T) {
	for i := 0; i < 10; i++ {
		cerateRandomEntry(t)
	}

	arg := ListEntryParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntry(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestDeleteEntry(t *testing.T) {
	entry1 := cerateRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
