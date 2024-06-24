package queries 

const GET_REPOSITORY_BY_NAME =`
	SELECT r.id, r.title, r.description, r.content, o.owner_id, u.name FROM REPOSITORY r
	JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id 
	JOIN COMMON_USER u ON o.owner_id = u.id WHERE r.title = ?`

const GET_REPOSITORIES = `
	SELECT r.id, r.title, r.description,r.content, o.owner_id, u.name FROM REPOSITORY r 
	JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id 
	JOIN COMMON_USER u ON o.owner_id = u.id `

const CREATE_REPOSITORY ="INSERT INTO REPOSITORY(title, description,content) VALUES (?,?,?)"

const UPDATE_CONTENT_REPOSITORY = `
	UPDATE REPOSITORY 
	SET content = ? 
	WHERE id = ?`
const UPDATE_TITLE_REPOSITORY = `
	UPDATE REPOSITORY 
	SET title = ? 
	WHERE id = ?`
const UPDATE_DESCRIPTION_REPOSITORY = `
	UPDATE REPOSITORY 
	SET description = ? 
	WHERE id = ?`

const GET_OWNER_NAME_BY_ID = "SELECT name FROM COMMON_USER WHERE id = ?"

const SET_OWNER = "INSERT INTO OWNERS_REPOSITORY(repository_id, owner_id) VALUES (?, ?)" 

const SET_CATEGORY = "INSERT INTO CATEGORY_REPO(category_id, repository_id) VALUES (?,?)"

const GET_CATEGORIES_BY_REPO = `
	SELECT c.name FROM CATEGORY_REPO cr
	JOIN CATEGORY c ON cr.category_id = c.id
	JOIN REPOSITORY r ON cr.repository_id = r.id
	WHERE cr.repository_id = ?
`
const GET_COMMENTS_BY_REPO = `
	SELECT com.id, u.id, u.name, com.comment FROM COMMENT com
	JOIN REPOSITORY r ON com.repository_id = r.id
	JOIN COMMON_USER u ON com.owner_id = u.id
	WHERE com.repository_id = ?
`
const SET_COMMENT = "INSERT INTO COMMENT (owner_id, repository_id, comment) VALUES (?, ?, ?)"