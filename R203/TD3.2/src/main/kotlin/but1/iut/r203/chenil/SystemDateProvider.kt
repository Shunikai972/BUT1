package but1.iut.r203.chenil

import java.time.LocalDate

/**
 * Implementation concrète de DateProvider utilisant l'horloge système
 * Cette classe utilise LocalDate.now() pour obtenir la date actuelle
 */
class SystemDateProvider : DateProvider {
    override fun getDate(): LocalDate {
        return LocalDate.now()
    }
}
