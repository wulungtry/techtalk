package common

import "fmt"

//ErrorDBConn return message for failed to connect to database
func ErrorDBConn() string {
	return "Koneksi database gagal"
}

//Exception Print default exception message
func Exception() string {
	return "Terjadi kesalahan. Mohon dicoba beberapa saat lagi."
}

//Fail to start either gRPC or REST service
func FailToStartServer(t string) string {
	return fmt.Sprintf("Start service %s gagal", t)
}

//Found Print default found message
func Found(module string) string {
	return fmt.Sprintf("Data %s berhasil ditemukan.", module)
}

//FoundCount Print default found message with the amount of data found
func FoundCount(count int, module string) string {
	return fmt.Sprintf("%d data %s berhasil ditemukan.", count, module)
}

//NotFound Print default not found message
func NotFound(module string) string {
	return fmt.Sprintf("Data %s tidak ditemukan.", module)
}

//InsertSuccess Print default succes message for insert operation
func InsertSuccess(module string) string {
	return fmt.Sprintf("Data %s berhasil disimpan.", module)
}

//InsertFail Print default fail message for insert operation
func InsertFail(module string) string {
	return fmt.Sprintf("Data %s gagal disimpan.", module)
}

//UpdateSuccess Print default succes message for update operation
func UpdateSuccess(module string) string {
	return fmt.Sprintf("Data %s berhasil diubah.", module)
}

//UpdateFail Print default fail message for update operation
func UpdateFail(module string) string {
	return fmt.Sprintf("Data %s gagal diubah.", module)
}

//DeleteSuccess Print default succes message for delete operation
func DeleteSuccess(module string) string {
	return fmt.Sprintf("Data %s berhasil dihapus.", module)
}

//DeleteFail Print default fail message for delete operation
func DeleteFail(module string) string {
	return fmt.Sprintf("Data %s gagal dihapus.", module)
}

//FileNotFound Print default file not found
func FileNotFound(file string) string {
	return fmt.Sprintf("File %s tidak ditemukan.", file)
}

//EmptyRequest Print default empty request
func EmptyRequest(data string) string {
	return fmt.Sprintf("Data request %s tidak tidak boleh kosong.", data)
}

//InvalidRequest Print default invalid request
func InvalidRequest(data string) string {
	return fmt.Sprintf("Data request %s tidak valid.", data)
}

//StartRESTServer
func StartRESTServer() string {
	return "Starting REST gateway server"
}

func StartGRPCServer() string {
	return "Starting GRPC gateway server"
}
