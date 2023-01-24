package mysql

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceRDS(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	binlog := acctest.RandIntRange(0, 78)
	targetDelay := acctest.RandIntRange(0, 7200)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRDSConfig_basic(rName, binlog, targetDelay),
				Check:  testAccRDSConfigExists(fmt.Sprintf("mysql_rds_config.%s", rName)),
			},
			{
				Config: testAccRDSConfig_basic(rName, binlog, targetDelay),
				Check:  resource.TestCheckResourceAttr(fmt.Sprintf("mysql_rds_config.%s", rName), "binlog_retention_period", fmt.Sprintf("%d", binlog)),
			},
		},
	})
}

func testAccRDSConfig_basic(rName string, binlog int, replication int) string {
	return fmt.Sprintf(`
resource "mysql_rds_config" "%s" {
                binlog_retention_period = %d
                replication_target_delay = %d
}`, rName, binlog, replication)
}

func testAccRDSConfigExists(rn string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("RDS config id not set")
		}

		return nil
	}
}
