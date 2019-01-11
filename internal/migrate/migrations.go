package migrate

import (
	"github.com/rubenv/sql-migrate"

	"github.com/hvs-fasya/psychbal/internal/utils"
)

func getSource() (migrations *migrate.MemoryMigrationSource) {
	var h string
	h, _ = utils.HashAndSalt([]byte("12345678"))
	migrations = &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id: "1",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS roles(
									id bigserial not null,
									role text not null,
									primary key(id),
									unique (role)
								);
					INSERT INTO roles (role) VALUES ('admin');
					INSERT INTO roles (role) VALUES ('client');`,
				},
				Down: []string{"DROP TABLE roles;"},
			},
			&migrate.Migration{
				Id: "2",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS users(
									id bigserial not null,
									email text not null,
									phone text not null,
									nickname text not null,
									name text not null,
									lastname text not null,
									role_id int not null,
									pswd_hash text not null,
									primary key(id),
									unique (email),
									unique (nickname),
									CONSTRAINT users_role_id_fkey foreign key (role_id) REFERENCES roles(id) ON DELETE CASCADE
								);
					INSERT INTO users (nickname, name, role_id, pswd_hash, lastname, phone, email) VALUES ('Nina', 'Nina',
						(SELECT id FROM roles WHERE role='` + utils.AdminRoleName + `'), '` + h + `', '', '', 'nina@example.com');
					INSERT INTO users (nickname, name, role_id, pswd_hash, lastname, phone, email) VALUES ('admin', 'admin',
						(SELECT id FROM roles WHERE role='` + utils.AdminRoleName + `'), '` + h + `', '', '', 'admin@example.com');`,
				},
				Down: []string{"ALTER TABLE users DROP CONSTRAINT users_role_id_fkey; DROP TABLE users;"},
			},
			&migrate.Migration{
				Id: "3",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS posts(
									id bigserial not null,
									title text not null,
									content text,
									published_at timestamp not null,
									deleted_at timestamp,
									primary key(id),
									unique (title))`,
				},
				Down: []string{"DROP TABLE posts"},
			},
			&migrate.Migration{
				Id: "4",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS comments(
									id bigserial not null,
									post_id int not null,
									comment_id int,
									user_id int not null,
									content text,
									hidden bool default false,
									primary key(id),
									CONSTRAINT comments_post_id_fkey foreign key (post_id) REFERENCES posts(id) ON DELETE CASCADE,
									CONSTRAINT comments_comment_id_fkey foreign key (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
									CONSTRAINT comments_user_id_fkey foreign key (user_id) REFERENCES users(id) ON DELETE CASCADE
								)`,
				},
				Down: []string{"ALTER TABLE comments DROP CONSTRAINT comments_user_id_fkey; " +
					"ALTER TABLE comments DROP CONSTRAINT comments_comment_id_fkey; " +
					"ALTER TABLE comments DROP CONSTRAINT comments_post_id_fkey; " +
					"DROP TABLE comments;"},
			},
		},
	}
	return
}
