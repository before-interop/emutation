# emutation
- v0.1: Dépôt du wsdl Initial
- v0.2: Ajout d'une version modifiée incluant les modifications suivantes
  - Changement du type "Char" en "String" pour le champs "ComplementNumeroVoie"
     - Le type char faisait que n'importe quelle donnée qui était retournée avait la valeur 0
  - Suppession des ArrayOf... car tout les tableaux sont gérés via des séquences
     - Les "ArrayOf" faisait que nous ne pouvions pas retourner un tableau avec un seul élément
  - Suppression des "maxLength" sur les <xs:element>
     - Le wsdl n'était pas valide avec ces attributs
- v0.3: Ajout d'une version modifiée incluant les modifications suivante
  - Modification des types "unsignedInt" en "nonNegativeInteger" 
     - Les valeurs étaient mal interprétées avec le type "unsignedInt"
  - Ajout d'un attribut sur le ListeFibreType
     - Ajout d'un maxOccurs="unbounded" car sans cet attribut on ne pouvait mette qu'un élément dans le tableau
  - Remplacement du type complexe "EtatFibreType" en énumérateur
     - Cela permettra que seule les valeurs prévues dans la documentation puisse passer
  
