# def cekpalindrom (teks) :
#     temp = []
#     belakang = ''
#     depan = ''
#     indikator = True
#     for i in range (len(teks)) :
#         if teks[i] != ',' and teks[i] != '.' and teks[i] != "'" and teks[i] != '"' and teks[i] != '!' and teks[i] != '?' and teks[i] != ' ' :
#             temp.append(teks[i])
#     while len(temp) > 1 and indikator  :
#         depan += temp.pop(0)
#         belakang += temp.pop()
#         if belakang != depan :
#             indikator = False
#     if indikator :
#         print(teks, "adalah kata palyndrom.", depan, belakang)
#     else :
#         print(teks, "bukan kata palyndrom.", indikator)


# a = input('masukkan teks : ')

# print('hasil : ')
# cekpalindrom(a)



a = input("masukan text: ")
s = a.split()[::-1]
l = []
for i in s:
    l.append(i)
    
print(" ".join(l))