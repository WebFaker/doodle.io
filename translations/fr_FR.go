package translations

func initFrenchTranslation() Translation {
	translation := createTranslation()

	translation.put("requires-js", "Ce site internet nécessite JavaScript pour fonctionner correctement.")

	translation.put("start-the-game", "Lancer le jeu")
	translation.put("start", "Commencer")
	translation.put("game-not-started-title", "Le jeu n'a pas encore commencé")
	translation.put("waiting-for-host-to-start", "En attente de l'hôte pour lancer la partie.")

	translation.put("round", "Manche")
	translation.put("toggle-soundeffects", "Effets sonores")
	translation.put("change-your-name", "Modifier votre pseudo")
	translation.put("randomize", "Aléatoire")
	translation.put("apply", "Appliquer")
	translation.put("save", "Sauvegarder")
	translation.put("toggle-fullscreen", "Plein écran")
	translation.put("show-help", "Aide")
	translation.put("votekick-a-player", "Exclure un joueur")

	translation.put("last-turn", "(Dernier tour : %s pts)")

	translation.put("drawer-kicked", "Le joueur exclu étant celui qui dessinait, personne n'emporte de points pour ce tour.")
	translation.put("self-kicked", "Vous avez été exclu")
	translation.put("kick-vote", "(%s/%s) joueurs ont voté pour exclure %s.")
	translation.put("player-kicked", "Le Joueur a été exclu.")
	translation.put("owner-change", "%s est le nouvel hôte de la partie.")

	translation.put("change-lobby-settings-tooltip", "Changer les paramètres de la partie")
	translation.put("change-lobby-settings-title", "Paramètres de la partie")
	translation.put("lobby-settings-changed", "Paramètres de la partie modifiés")
	translation.put("advanced-settings", "Paramètres avancés")
	translation.put("word-language", "Langue des mots")
	translation.put("drawing-time-setting", "Temps de dessin")
	translation.put("rounds-setting", "Manches")
	translation.put("max-players-setting", "Nb de joueurs max.")
	translation.put("public-lobby-setting", "Partie publique")
	translation.put("custom-words", "Mots personnalisés")
	translation.put("custom-words-info", "Écrivez vos mots, séparés par des virgules")
	translation.put("custom-words-chance-setting", "Taux d'apparition des mots perso.")
	translation.put("players-per-ip-limit-setting", "Nb de joueurs max par IP")
	translation.put("enable-votekick-setting", "Autoriser l'exclusion")
	translation.put("save-settings", "Sauvegarder les paramètres")
	translation.put("input-contains-invalid-data", "Le champs contient des données invalides :")
	translation.put("please-fix-invalid-input", "Corrigez le champs requis et réessayez.")
	translation.put("create-lobby", "Créer une partie")

	translation.put("players", "Joueurs")
	translation.put("refresh", "Rafraîchir")
	translation.put("join-lobby", "Rejoindre la partie")

	translation.put("message-input-placeholder", "Écrivez vos réponses et messages ici")

	translation.put("choose-a-word", "Choisissez un mot")
	translation.put("waiting-for-word-selection", "En attente du mot choisi")
	//This one doesn't use %s, since we want to make one part bold.
	translation.put("is-choosing-word", "est en train de choisir.")

	translation.put("close-guess", "'%s' est très proche.")
	translation.put("correct-guess", "Vous avez deviné le mot.")
	translation.put("correct-guess-other-player", "'%s' a deviné le mot.")
	translation.put("round-over", "Tour terminé, aucun mot n'a été choisi.")
	translation.put("round-over-no-word", "Tour terminé, le mot était '%s'.")
	translation.put("game-over-win", "Bravo, vous avez gagné !")
	translation.put("game-over-tie", "Égalité !")
	translation.put("game-over", "Vous êtes %se avec %s points")

	translation.put("change-active-color", "Changer la couleur")
	translation.put("use-pencil", "Crayon")
	translation.put("use-eraser", "Gomme")
	translation.put("use-fill-bucket", "Remplissage (Remplis la zone ciblée avec la couleur sélectionée)")
	translation.put("change-pencil-size-to", "Changer la taille du crayon / de la gomme vers %s")
	translation.put("clear-canvas", "Tout effacer")
	translation.put("undo", "Annuler le dernier changement (Ne fonctionne pas après un \""+translation.Get("clear-canvas")+"\")")

	translation.put("connection-lost", "Connection perdue !")
	translation.put("connection-lost-text", "Tentative de reconnexion en cours..."+
		" ...\n\nVérifiez votre connexion internet.\nSi le "+
		"problème persiste, contactez-nous.")
	translation.put("error-connecting", "Erreur de connexion au serveur")
	translation.put("error-connecting-text",
		"Doodle.io n'a pas pu établir une connexion socket.\n\nVotre connexion internet "+
			"semble fonctionner, mais le\nserveur ou votre antivirus "+
			"n'a peut-être pas été configuré correctement.\n\nPour réessayer, rechargez la page.")

	translation.put("message-too-long", "Votre message est trop long.")

	//Help dialog
	translation.put("controls", "Contrôles")
	translation.put("pencil", "Crayon")
	translation.put("eraser", "Gomme")
	translation.put("fill-bucket", "Remplissage")
	translation.put("switch-tools-intro", "Vous pouvez changer d'outils en utilisant les raccourcis")
	translation.put("switch-pencil-sizes", "Vous pouvez changer la taille du crayon / de la gomme en appuyant sur %s à %s.")

	//Generic words
	//"close" as in "closing the window"
	translation.put("close", "Fermer")
	translation.put("no", "Non")
	translation.put("yes", "Oui")
	translation.put("system", "Système")

	translation.put("source-code", "Code Source")
	translation.put("help", "Aide")
	translation.put("contact", "Contact")
	translation.put("submit-feedback", "Retours")
	translation.put("stats", "Statut")

	RegisterTranslation("fr", translation)
	RegisterTranslation("fr-fr", translation)

	return translation
}
