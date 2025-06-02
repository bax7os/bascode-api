package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios creates a new repository of users
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar inserts a user in the database
func (u *usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := u.db.Prepare("insert into usuarios (nome, nick, email, senha, foto_perfil) values (?, ?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha, usuario.FotoPerfil)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar gets all users with name or nick
func (repositorio *usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick,
		nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEM,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (repositorios usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorios.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEM,
			&usuario.FotoPerfil,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
		return usuario, nil
	}

	return modelos.Usuario{}, nil
}

func (repositorios usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorios.db.Prepare("update usuarios set nome = ?, nick = ?, email = ?, foto_perfil = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.FotoPerfil, ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorios usuarios) Deletar(ID uint64) error {
	statement, erro := repositorios.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorios usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorios.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()
	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}

	}
	return usuario, nil
}

func (repositorios usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorios.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	fmt.Println(usuarioID)
	fmt.Println(seguidorID)

	return nil

}

func (repositorios usuarios) DeixarDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorios.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

func (repositorios usuarios) BuscarSeguidores(usuario_id uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorios.db.Query(
		"select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?",
		usuario_id,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEM,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil

}

func (repositorios usuarios) BuscarSeguindo(usuario_id uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorios.db.Query(
		"select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u  inner join seguidores s on u.id = s.usuario_id where seguidor_id = ?",
		usuario_id,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEM,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil

}

func (repositorios usuarios) BuscarSenha(usuario_id uint64) (string, error) {
	linha, erro := repositorios.db.Query("select senha from usuarios where id = ?", usuario_id)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()
	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

func (repositorios usuarios) AtualizarSenha(usuario_id uint64, senha string) error {
	statement, erro := repositorios.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(senha, usuario_id); erro != nil {
		return erro
	}
	return nil
}

func (repositorios usuarios) AtualizarFotoPerfil(usuarioID uint64, caminhoArquivo string) error {
	statement, erro := repositorios.db.Prepare("update usuarios set foto_perfil = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(caminhoArquivo, usuarioID); erro != nil {
		return erro
	}
	return nil
}
