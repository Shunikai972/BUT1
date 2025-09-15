    bool estBissextile(unsigned int a)
    {
        if (a >= 1582)
        {
            if ((a % 400 == 0) || (a % 4 == 0 && a % 100 != 0))
                  return(true); 
             else 
                  return(false);
        }

        else 
        {
            return (false);
        }
    }
    short joursDansMois(unsigned int annee, unsigned short mois)
    {
            if (annee < 1582 || mois < 1 || mois > 12)
            {
                return (-1);
            }
            else if  (mois == 1 || mois == 3 || mois == 5 || mois == 7 || mois == 8 || mois == 10 || mois == 12)
            {
                return(31);
            }
            else if  (mois == 4 || mois == 6 || mois == 9 || mois == 11)
            {
                return(30);
            }
            else if (mois == 2)
            {
                if (estBissextile(annee) == true)
                    return (29);
                else 
                    return (28);
            }
            else
                return(-1);
    }

    bool estJourValide(unsigned short jour, unsigned short mois,unsigned int annee)
    {
        if (annee < 1582 || mois <  1 || mois > 12 || jour < 1 || jour > joursDansMois(annee, mois))
            return (false);
        else
            return(true);
    }

    int joursDepuisOrigine(unsigned short jour, unsigned short mois,unsigned int annee)
    {
        int nbJours = 0;
        for (unsigned int anneeCourante=1582; anneeCourante <annee;anneeCourante++)
            nbJours += estBissextile(anneeCourante) ? 366 : 365;
        for (unsigned int moisCourant=1; moisCourant< mois; moisCourant++)
            nbJours += joursDansMois(annee,moisCourant);
        nbJours += jour - 1;
        return (nbJours);
    }
    
    int jourDepuisNaissanceTuring(unsigned short jour, unsigned short mois,unsigned int annee)
    {
        int a;
        int b;
        int temp;
        if (estJourValide(jour,mois, annee)) {
            a = joursDepuisOrigine(23,6,1912);
            b = joursDepuisOrigine(jour,mois,annee);
        
            if(a > b)
                return(a-b);
            else {
               temp = a - b;
               temp = -temp;
               return (temp);
            }
        }
        return(-1);
    }