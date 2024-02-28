game_id := $(shell cat /tmp/game_id.txt 2>/dev/null || echo 0)
salt := $(shell cat /tmp/salt.txt 2>/dev/null || echo 0)

mock-expected-keepers:
	mockgen -source=x/whichnumber/types/expected_keepers.go \
		-package testutil \
		-destination=x/whichnumber/testutil/expected_keepers_mocks.go

new-game:
	@echo "Creating a new game"
	@OUTPUT=$$(whichnumberd tx whichnumber new-game 42 100stake 20000stake --from bob --yes) && \
	echo "$${OUTPUT}" && \
	game_id=$$(echo "$${OUTPUT}" | grep -oE 'game_id.*[0-9]+' | grep -oE '[0-9]+' | head -n 1) && \
    echo "New game created with ID: $${game_id}" && \
	echo "$${game_id}" > /tmp/game_id.txt

commit-number:
	@echo "Committing for game $(game_id)"
	@OUTPUT=$$(whichnumberd tx whichnumber commit-number $(game_id) 38 --from alice --yes) && \
	echo "$${OUTPUT}" && \
	salt=$$(echo "$${OUTPUT}" | grep '^copy your salt for the reveal:  ' | sed 's/^copy your salt for the reveal:  //') && \
	echo "Salt: $${salt}" && \
	echo "$${salt}" > /tmp/salt.txt

reveal-number:
	@echo "Revealing for game $(game_id)"
	whichnumberd tx whichnumber reveal-number $(game_id) 38 $(salt) --from alice --yes

show-game:
	@echo "Showing game $(game_id)"
	whichnumberd query whichnumber show-game $(game_id)

list-games:
	@echo "Listing games"
	whichnumberd query whichnumber list-games

commit_timeout := 120
reveal_timeout := 120
max_players := 1
min_distance_to_win := 12

update-params:
	@echo "Updating params"
	whichnumberd tx whichnumber update-params ${commit_timeout} ${reveal_timeout} ${max_players} ${min_distance_to_win} --from bob --yes

show-params:
	@echo "Showing params"
	whichnumberd query whichnumber params

show-system:
	@echo "Showing system"
	whichnumberd query whichnumber show-system-info
