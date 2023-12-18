func main() {
	rootCmd, _ := pyloncmd.NewRootCmd()
	rootCmd.Short = "Stargate Pylons App"
	rootCmd.AddCommand(pyloncmd.DevValidate())
      "@templates": ["./@next/components/templates"],
      "@pages": ["./@next/pages"],
      "@layouts": ["./@next/layouts"]
    }
  },

func (suite *IntegrationTestSuite) TestAppleIAPOrderExist() {
	k := suite.k
	ctx := suite.ctx
	require := suite.Require()

	items := createNAppleIAPOrder(k, ctx, 10)
	for _, item := range items {
		require.True(k.HasAppleIAPOrder(ctx, item.PurchaseId))
	}
}
