package mysql

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceRDS(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRDSConfig_basic(10, 30),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("mysql_rds_config.test", "binlog_retention_period", "10"),
					resource.TestCheckResourceAttr("mysql_rds_config.test", "replication_target_delay", "30"),
				),
			},
		},
	})
}

func testAccRDSConfig_basic(binlog int, replication int) string {
	return fmt.Sprintf(`
resource "mysql_rds_config" "test" {
                binlog_retention_period = %d
                replication_target_delay = %d
}`, binlog, replication)
}
