package but1.iut.r203.chenil

import java.time.LocalDate

class DateConsole {
    fun getDate(): LocalDate {
        print("Année?")
        val anneeConsole = Integer.valueOf(readLine())
        print("Mois ?")
        val moisConsole = Integer.valueOf(readLine())
        print("Jour ?")
        val jourConsole = Integer.valueOf(readLine())
        return LocalDate.of(anneeConsole,moisConsole, jourConsole)
    }
}