module github.com/reifiedbeans/project-zombot

go 1.19

// Remove me once https://github.com/Amatsagu/Tempest/commit/0b2d29a6b73f7f80283d7e229edb2d3a21bad75c is released
replace github.com/Amatsagu/Tempest => github.com/reifiedbeans/Tempest v0.0.0-20230124113543-0b2d29a6b73f

require (
	github.com/Amatsagu/Tempest v1.0.3
	github.com/google/uuid v1.3.0
	github.com/gorcon/rcon v1.3.4
	github.com/pkg/errors v0.8.1
	go.uber.org/config v1.4.0
	go.uber.org/fx v1.19.1
	go.uber.org/zap v1.24.0
)

require (
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/dig v1.16.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/lint v0.0.0-20190930215403-16217165b5de // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/tools v0.5.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
