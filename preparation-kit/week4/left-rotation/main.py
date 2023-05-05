
def rotate_left(d: int, arr: list) -> list:
    """
    Rota un arreglo a la izquierda d veces

    :param d: numeoro de rotaciones
    :param arr: arreglo a rotar
    :return: el arreglo rotado
    """

    for _ in range(d):
        arr.append(arr.pop(0))

    return arr

def main():
    d_test=2
    arr_test=[1,2,3,4,5]

    arr = rotate_left(d_test, arr_test)

    print(arr)

if __name__ == '__main__':
    main()
