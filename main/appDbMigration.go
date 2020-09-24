package main

import (
	"fmt"
	"myfirstgosql/config"
)

func dbMigration(sf *config.SessionFactory) {
	session := sf.GetSession()

	fmt.Printf("%-30s %30s\n", "Creating M_CATEGORY table", "[OK]")
	_, err := session.Exec("CREATE TABLE IF NOT EXISTS `m_category` (`id` char(36) NOT NULL,`category_name` varchar(255) DEFAULT NULL,`created_at` datetime DEFAULT CURRENT_TIMESTAMP,`updated_at` datetime DEFAULT CURRENT_TIMESTAMP,`deleted_at` datetime DEFAULT NULL,PRIMARY KEY (`id`))")
	if err != nil {
		fmt.Printf("%-30s %30s\n", "Creating M_CATEGORY table", "[FAILED]")
		panic(err)
	}

	fmt.Printf("%-30s %30s\n", "Creating M_PRODUCT table", "[OK]")
	_, err = session.Exec("CREATE TABLE IF NOT EXISTS `m_product` (`id` char(36) NOT NULL,`product_code` varchar(255) NOT NULL,`product_name` varchar(255) NOT NULL,`created_at` datetime DEFAULT CURRENT_TIMESTAMP,`updated_at` datetime DEFAULT CURRENT_TIMESTAMP,`deleted_at` datetime DEFAULT NULL,`category_id` char(36) DEFAULT NULL,PRIMARY KEY (`id`),UNIQUE KEY `product_code_UNIQUE` (`product_code`),KEY `category_id` (`category_id`),CONSTRAINT `m_product_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `m_category` (`id`) ON DELETE SET NULL ON UPDATE CASCADE)")
	if err != nil {
		fmt.Printf("%-30s %30s\n", "Creating M_PRODUCT table", "[FAILED]")
		panic(err)
	}

	fmt.Printf("%-30s %30s\n", "Creating M_PRODUCT_PRICE table", "[OK]")
	_, err = session.Exec("CREATE TABLE IF NOT EXISTS `m_product_price` (`product_price_id` char(36) NOT NULL,`product_id` char(36) NOT NULL,`product_price` int NOT NULL DEFAULT '0',`is_active` varchar(1) DEFAULT '0',PRIMARY KEY (`product_price_id`))")
	if err != nil {
		fmt.Printf("%-30s %30s\n", "Creating M_PRODUCT_PRICE table", "[FAILED]")
		panic(err)
	}
}
