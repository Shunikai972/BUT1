package ihm.td4.heureMVC.modele


import java.text.SimpleDateFormat
import java.util.*


class Model {
    var heure: String=""


    fun getTime(): String {

        val df = SimpleDateFormat("HH:mm:ss")
        this.heure=df.format(Calendar.getInstance().time)
        return this.heure

    }





}